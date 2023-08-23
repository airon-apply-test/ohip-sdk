/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Room Configuration API
 * APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { BedTypeType } from './BedTypeType';
import {
    BedTypeTypeFromJSON,
    BedTypeTypeFromJSONTyped,
    BedTypeTypeToJSON,
} from './BedTypeType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Response object for fetching Bed Types.
 * @export
 * @interface BedTypesDetails
 */
export interface BedTypesDetails {
    /**
     * Bed Type Enumeration element.
     * @type {Array<BedTypeType>}
     * @memberof BedTypesDetails
     */
    bedTypes?: Array<BedTypeType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof BedTypesDetails
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof BedTypesDetails
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the BedTypesDetails interface.
 */
export function instanceOfBedTypesDetails(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BedTypesDetailsFromJSON(json: any): BedTypesDetails {
    return BedTypesDetailsFromJSONTyped(json, false);
}

export function BedTypesDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): BedTypesDetails {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bedTypes': !exists(json, 'bedTypes') ? undefined : ((json['bedTypes'] as Array<any>).map(BedTypeTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function BedTypesDetailsToJSON(value?: BedTypesDetails | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bedTypes': value.bedTypes === undefined ? undefined : ((value.bedTypes as Array<any>).map(BedTypeTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

