/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block Configuration API
 * APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
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
import type { SalesManagerGoalType } from './SalesManagerGoalType';
import {
    SalesManagerGoalTypeFromJSON,
    SalesManagerGoalTypeFromJSONTyped,
    SalesManagerGoalTypeToJSON,
} from './SalesManagerGoalType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * You can use this API to retrieve individual Sales Manager Goals for a hotel, you can narrow the results using different search criteria
 * @export
 * @interface SalesManagerGoalsInfo
 */
export interface SalesManagerGoalsInfo {
    /**
     * Detail Information about Sales Manager's goal.
     * @type {Array<SalesManagerGoalType>}
     * @memberof SalesManagerGoalsInfo
     */
    salesManagerGoals?: Array<SalesManagerGoalType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof SalesManagerGoalsInfo
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof SalesManagerGoalsInfo
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the SalesManagerGoalsInfo interface.
 */
export function instanceOfSalesManagerGoalsInfo(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SalesManagerGoalsInfoFromJSON(json: any): SalesManagerGoalsInfo {
    return SalesManagerGoalsInfoFromJSONTyped(json, false);
}

export function SalesManagerGoalsInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): SalesManagerGoalsInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'salesManagerGoals': !exists(json, 'salesManagerGoals') ? undefined : ((json['salesManagerGoals'] as Array<any>).map(SalesManagerGoalTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function SalesManagerGoalsInfoToJSON(value?: SalesManagerGoalsInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'salesManagerGoals': value.salesManagerGoals === undefined ? undefined : ((value.salesManagerGoals as Array<any>).map(SalesManagerGoalTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

