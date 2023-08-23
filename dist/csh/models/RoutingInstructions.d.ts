/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { InstanceLink } from './InstanceLink';
import type { ResvRoutingCriteriaType } from './ResvRoutingCriteriaType';
import type { WarningType } from './WarningType';
/**
 * Request to create a routing instruction.
 * @export
 * @interface RoutingInstructions
 */
export interface RoutingInstructions {
    /**
     *
     * @type {ResvRoutingCriteriaType}
     * @memberof RoutingInstructions
     */
    criteria?: ResvRoutingCriteriaType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof RoutingInstructions
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof RoutingInstructions
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the RoutingInstructions interface.
 */
export declare function instanceOfRoutingInstructions(value: object): boolean;
export declare function RoutingInstructionsFromJSON(json: any): RoutingInstructions;
export declare function RoutingInstructionsFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoutingInstructions;
export declare function RoutingInstructionsToJSON(value?: RoutingInstructions | null): any;
