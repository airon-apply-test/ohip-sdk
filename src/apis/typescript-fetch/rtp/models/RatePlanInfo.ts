/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Rate API
 * APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { MasterInfoType } from './MasterInfoType';
import {
    MasterInfoTypeFromJSON,
    MasterInfoTypeFromJSONTyped,
    MasterInfoTypeToJSON,
} from './MasterInfoType';
import type { RatePlanType } from './RatePlanType';
import {
    RatePlanTypeFromJSON,
    RatePlanTypeFromJSONTyped,
    RatePlanTypeToJSON,
} from './RatePlanType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Response object to fetch rate plan's complete details.
 * @export
 * @interface RatePlanInfo
 */
export interface RatePlanInfo {
    /**
     * Rate plan complete details such as primary details, classifications, room types etc.
     * @type {Array<RatePlanType>}
     * @memberof RatePlanInfo
     */
    ratePlans?: Array<RatePlanType>;
    /**
     * Refer to Generic common types document.
     * @type {Array<MasterInfoType>}
     * @memberof RatePlanInfo
     */
    masterInfoList?: Array<MasterInfoType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof RatePlanInfo
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof RatePlanInfo
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the RatePlanInfo interface.
 */
export function instanceOfRatePlanInfo(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RatePlanInfoFromJSON(json: any): RatePlanInfo {
    return RatePlanInfoFromJSONTyped(json, false);
}

export function RatePlanInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): RatePlanInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'ratePlans': !exists(json, 'ratePlans') ? undefined : ((json['ratePlans'] as Array<any>).map(RatePlanTypeFromJSON)),
        'masterInfoList': !exists(json, 'masterInfoList') ? undefined : ((json['masterInfoList'] as Array<any>).map(MasterInfoTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function RatePlanInfoToJSON(value?: RatePlanInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ratePlans': value.ratePlans === undefined ? undefined : ((value.ratePlans as Array<any>).map(RatePlanTypeToJSON)),
        'masterInfoList': value.masterInfoList === undefined ? undefined : ((value.masterInfoList as Array<any>).map(MasterInfoTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

