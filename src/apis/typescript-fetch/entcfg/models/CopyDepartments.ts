/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CopyConfigurationCodeType } from './CopyConfigurationCodeType';
import {
    CopyConfigurationCodeTypeFromJSON,
    CopyConfigurationCodeTypeFromJSONTyped,
    CopyConfigurationCodeTypeToJSON,
} from './CopyConfigurationCodeType';
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
 * 
 * @export
 * @interface CopyDepartments
 */
export interface CopyDepartments {
    /**
     * List of the departments to be copied.
     * @type {Array<CopyConfigurationCodeType>}
     * @memberof CopyDepartments
     */
    departments?: Array<CopyConfigurationCodeType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof CopyDepartments
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CopyDepartments
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the CopyDepartments interface.
 */
export function instanceOfCopyDepartments(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CopyDepartmentsFromJSON(json: any): CopyDepartments {
    return CopyDepartmentsFromJSONTyped(json, false);
}

export function CopyDepartmentsFromJSONTyped(json: any, ignoreDiscriminator: boolean): CopyDepartments {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'departments': !exists(json, 'departments') ? undefined : ((json['departments'] as Array<any>).map(CopyConfigurationCodeTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function CopyDepartmentsToJSON(value?: CopyDepartments | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'departments': value.departments === undefined ? undefined : ((value.departments as Array<any>).map(CopyConfigurationCodeTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

